package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/core"
	"github.com/ApurvaBardapurkar/sdk-warps-go/pkg/types"
)

// WarpValidator provides functionality for validating warps
type WarpValidator struct {
	config types.WarpConfig
	schema map[string]interface{}
}

// NewWarpValidator creates a new WarpValidator instance
func NewWarpValidator(config types.WarpConfig) *WarpValidator {
	return &WarpValidator{
		config: config,
		schema: nil,
	}
}

// Validate validates a warp against the schema
func (v *WarpValidator) Validate(warp *types.Warp) error {
	if warp == nil {
		return errors.New("WarpValidator: warp is nil")
	}

	// Basic validation - ensure required fields are present
	if warp.Protocol == "" {
		return errors.New("WarpValidator: protocol is required")
	}
	if warp.Name == "" {
		return errors.New("WarpValidator: name is required")
	}
	if warp.Title == "" {
		return errors.New("WarpValidator: title is required")
	}
	if len(warp.Actions) == 0 {
		return errors.New("WarpValidator: at least one action is required")
	}

	// Validate each action
	for i, action := range warp.Actions {
		if err := v.validateAction(action); err != nil {
			return fmt.Errorf("WarpValidator: invalid action at index %d: %w", i, err)
		}
	}

	// In a real implementation, you would validate against a JSON schema
	// For this example, we'll just do some basic validation
	
	return nil
}

// validateAction validates a warp action
func (v *WarpValidator) validateAction(action types.WarpAction) error {
	if action == nil {
		return errors.New("action is nil")
	}

	actionType := action.GetType()
	label := action.GetLabel()

	if actionType == "" {
		return errors.New("action type is required")
	}
	if label == "" {
		return errors.New("action label is required")
	}

	// Validate based on action type
	switch actionType {
	case types.TransferActionType:
		if a, ok := action.(types.WarpTransferAction); ok {
			return v.validateTransferAction(a)
		}
	case types.ContractActionType:
		if a, ok := action.(types.WarpContractAction); ok {
			return v.validateContractAction(a)
		}
	case types.QueryActionType:
		if a, ok := action.(types.WarpQueryAction); ok {
			return v.validateQueryAction(a)
		}
	case types.CollectActionType:
		if a, ok := action.(types.WarpCollectAction); ok {
			return v.validateCollectAction(a)
		}
	case types.LinkActionType:
		if a, ok := action.(types.WarpLinkAction); ok {
			return v.validateLinkAction(a)
		}
	default:
		return fmt.Errorf("unsupported action type: %s", actionType)
	}

	return nil
}

// validateTransferAction validates a transfer action
func (v *WarpValidator) validateTransferAction(action types.WarpTransferAction) error {
	// Validate required fields
	if action.Type != types.TransferActionType {
		return fmt.Errorf("expected action type %s, got %s", types.TransferActionType, action.Type)
	}

	// Validate inputs
	if action.Inputs != nil {
		for i, input := range action.Inputs {
			if err := v.validateInput(input); err != nil {
				return fmt.Errorf("invalid input at index %d: %w", i, err)
			}
		}
	}

	return nil
}

// validateContractAction validates a contract action
func (v *WarpValidator) validateContractAction(action types.WarpContractAction) error {
	// Validate required fields
	if action.Type != types.ContractActionType {
		return fmt.Errorf("expected action type %s, got %s", types.ContractActionType, action.Type)
	}
	if action.Address == "" {
		return errors.New("contract address is required")
	}

	// Validate inputs
	if action.Inputs != nil {
		for i, input := range action.Inputs {
			if err := v.validateInput(input); err != nil {
				return fmt.Errorf("invalid input at index %d: %w", i, err)
			}
		}
	}

	return nil
}

// validateQueryAction validates a query action
func (v *WarpValidator) validateQueryAction(action types.WarpQueryAction) error {
	// Validate required fields
	if action.Type != types.QueryActionType {
		return fmt.Errorf("expected action type %s, got %s", types.QueryActionType, action.Type)
	}
	if action.Address == "" {
		return errors.New("query address is required")
	}
	if action.Func == "" {
		return errors.New("query function is required")
	}

	// Validate inputs
	if action.Inputs != nil {
		for i, input := range action.Inputs {
			if err := v.validateInput(input); err != nil {
				return fmt.Errorf("invalid input at index %d: %w", i, err)
			}
		}
	}

	return nil
}

// validateCollectAction validates a collect action
func (v *WarpValidator) validateCollectAction(action types.WarpCollectAction) error {
	// Validate required fields
	if action.Type != types.CollectActionType {
		return fmt.Errorf("expected action type %s, got %s", types.CollectActionType, action.Type)
	}
	if action.Destination.URL == "" {
		return errors.New("destination URL is required")
	}

	// Validate URL
	_, err := url.Parse(action.Destination.URL)
	if err != nil {
		return fmt.Errorf("invalid destination URL: %w", err)
	}

	// Validate method
	if action.Destination.Method != types.GET && action.Destination.Method != types.POST {
		return fmt.Errorf("unsupported HTTP method: %s", action.Destination.Method)
	}

	// Validate inputs
	if action.Inputs != nil {
		for i, input := range action.Inputs {
			if err := v.validateInput(input); err != nil {
				return fmt.Errorf("invalid input at index %d: %w", i, err)
			}
		}
	}

	return nil
}

// validateLinkAction validates a link action
func (v *WarpValidator) validateLinkAction(action types.WarpLinkAction) error {
	// Validate required fields
	if action.Type != types.LinkActionType {
		return fmt.Errorf("expected action type %s, got %s", types.LinkActionType, action.Type)
	}
	if action.URL == "" {
		return errors.New("URL is required")
	}

	// Validate URL
	_, err := url.Parse(action.URL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	// Validate inputs
	if action.Inputs != nil {
		for i, input := range action.Inputs {
			if err := v.validateInput(input); err != nil {
				return fmt.Errorf("invalid input at index %d: %w", i, err)
			}
		}
	}

	return nil
}

// validateInput validates an input
func (v *WarpValidator) validateInput(input types.WarpActionInput) error {
	// Validate required fields
	if input.Name == "" {
		return errors.New("input name is required")
	}
	if input.Type == "" {
		return errors.New("input type is required")
	}
	if input.Position == "" {
		return errors.New("input position is required")
	}
	if input.Source == "" {
		return errors.New("input source is required")
	}

	// Validate source
	validSources := map[types.WarpActionInputSource]bool{
		types.FieldSource: true,
		types.QuerySource: true,
	}
	if !validSources[input.Source] {
		return fmt.Errorf("invalid input source: %s", input.Source)
	}

	// Validate position based on format
	if input.Position == types.ReceiverPosition || input.Position == types.ValuePosition || input.Position == types.TransferPosition {
		// These are valid static positions
	} else {
		// Check if it's an arg position (arg:1, arg:2, etc.)
		argPositionPattern := regexp.MustCompile(`^arg:[1-9][0-9]*$`)
		if !argPositionPattern.MatchString(string(input.Position)) {
			return fmt.Errorf("invalid input position: %s", input.Position)
		}
	}

	// Additional validation based on input type could be added here

	return nil
}

// loadSchema loads the schema from the specified URL
func (v *WarpValidator) loadSchema() error {
	if v.schema != nil {
		return nil
	}

	schemaURL := v.config.WarpSchemaURL
	if schemaURL == "" {
		schemaURL = core.Config.DefaultWarpSchemaURL(v.config.Env)
	}

	resp, err := http.Get(schemaURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to load schema: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(body, &schema); err != nil {
		return err
	}

	v.schema = schema
	return nil
} 