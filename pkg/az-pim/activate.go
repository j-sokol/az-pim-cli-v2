package azpim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ActivateRoleAssReq struct {
	Properties struct {
		PrincipalId                     string `json:"PrincipalId"`
		RoleDefinitionId                string `json:"RoleDefinitionId"`
		RequestType                     string `json:"RequestType"`
		LinkedRoleEligibilityScheduleId string `json:"LinkedRoleEligibilityScheduleId"`
		Justification                   string `json:"Justification"`
		ScheduleInfo                    struct {
			StartDateTime interface{} `json:"StartDateTime"`
			Expiration    struct {
				Duration string `json:"Duration"`
				Type     string `json:"Type"`
			} `json:"Expiration"`
		} `json:"ScheduleInfo"`
		TicketInfo struct {
			TicketNumber string `json:"TicketNumber"`
			TicketSystem string `json:"TicketSystem"`
		} `json:"TicketInfo"`
		IsValidationOnly bool `json:"IsValidationOnly"`
		IsActivativation bool `json:"IsActivativation"`
	} `json:"Properties"`
}

func ActivateRoleAssignment(token, userObjectId string, role Role) error {
	client := &http.Client{}
	ars := &ActivateRoleAssReq{
		Properties: struct {
			PrincipalId                     string `json:"PrincipalId"`
			RoleDefinitionId                string `json:"RoleDefinitionId"`
			RequestType                     string `json:"RequestType"`
			LinkedRoleEligibilityScheduleId string `json:"LinkedRoleEligibilityScheduleId"`
			Justification                   string `json:"Justification"`
			ScheduleInfo                    struct {
				StartDateTime interface{} `json:"StartDateTime"`
				Expiration    struct {
					Duration string `json:"Duration"`
					Type     string `json:"Type"`
				} `json:"Expiration"`
			} `json:"ScheduleInfo"`
			TicketInfo struct {
				TicketNumber string `json:"TicketNumber"`
				TicketSystem string `json:"TicketSystem"`
			} `json:"TicketInfo"`
			IsValidationOnly bool `json:"IsValidationOnly"`
			IsActivativation bool `json:"IsActivativation"`
		}{
			PrincipalId:                     userObjectId,
			RoleDefinitionId:                role.Properties.RoleDefinitionId,
			RequestType:                     "SelfActivate",
			LinkedRoleEligibilityScheduleId: role.Properties.RoleEligibilityScheduleId,
			Justification:                   "role perm",
			ScheduleInfo: struct {
				StartDateTime interface{} `json:"StartDateTime"`
				Expiration    struct {
					Duration string `json:"Duration"`
					Type     string `json:"Type"`
				} `json:"Expiration"`
			}{
				StartDateTime: nil,
				Expiration: struct {
					Duration string `json:"Duration"`
					Type     string `json:"Type"`
				}{
					Duration: "PT480M",
					Type:     "AfterDuration",
				},
			},
			TicketInfo: struct {
				TicketNumber string `json:"TicketNumber"`
				TicketSystem string `json:"TicketSystem"`
			}{
				TicketNumber: "",
				TicketSystem: "",
			},
			IsValidationOnly: false,
			IsActivativation: true,
		},
	}

	arsJson, err := json.Marshal(ars)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}
	guid := uuid.New()

	url := "https://management.azure.com" + role.Properties.Scope + "/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/" + guid.String() + "?api-version=2020-10-01"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(arsJson)))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to activate role: %v", resp.StatusCode)
	}

	return nil
}
