package azpim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RoleEligibilityResp struct {
	Role []Role `json:"value"`
}

type Role struct {
	Properties struct {
		RoleEligibilityScheduleId string `json:"roleEligibilityScheduleId"`
		Scope                     string `json:"scope"`
		RoleDefinitionId          string `json:"roleDefinitionId"`
		PrincipalId               string `json:"principalId"`
		PrincipalType             string `json:"principalType"`
		Status                    string `json:"status"`
		StartDateTime             string `json:"startDateTime"`
		MemberType                string `json:"memberType"`
		CreatedOn                 string `json:"createdOn"`
		ExpandedProperties        struct {
			Principal struct {
				ID          string `json:"id"`
				DisplayName string `json:"displayName"`
				Type        string `json:"type"`
			} `json:"principal"`
			RoleDefinition struct {
				ID          string `json:"id"`
				DisplayName string `json:"displayName"`
				Type        string `json:"type"`
			} `json:"roleDefinition"`
			Scope struct {
				ID          string `json:"id"`
				DisplayName string `json:"displayName"`
				Type        string `json:"type"`
			} `json:"scope"`
		} `json:"expandedProperties"`
	} `json:"properties"`
	Name string `json:"name"`
	ID   string `json:"id"`
	Type string `json:"type"`
}

func ListEligibleRoleAssignments(token string) (RoleEligibilityResp, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://management.azure.com/providers/Microsoft.Authorization/roleEligibilityScheduleInstances?api-version=2020-10-01&$filter=asTarget()", nil)
	if err != nil {
		return RoleEligibilityResp{}, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return RoleEligibilityResp{}, err
	}
	defer resp.Body.Close()
	var result RoleEligibilityResp
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

func GetEligibleRoleAssignment(token, scope, role string) (Role, error) {
	ra, err := ListEligibleRoleAssignments(token)
	if err != nil {
		return Role{}, fmt.Errorf("failed to list eligible role assignments: %v", err)
	}
	for _, v := range ra.Role {
		if strings.ToLower(v.Properties.ExpandedProperties.Scope.DisplayName) == strings.ToLower(scope) &&
			strings.ToLower(v.Properties.ExpandedProperties.RoleDefinition.DisplayName) == strings.ToLower(role) {
			return v, nil
		}
	}
	return Role{}, fmt.Errorf("role %s on %s not found", role, scope)
}
