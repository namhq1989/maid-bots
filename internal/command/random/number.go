package random

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/namhq1989/maid-bots/util/appcommand"

	"github.com/namhq1989/maid-bots/pkg/sentryio"

	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Number struct {
	Message    string
	Parameters *NumberParameters
}

type NumberParameters struct {
	Kind   string
	Min    int
	Max    int
	Count  int
	Unique bool
}

var (
	kindInt    = "int"
	kindDouble = "double"
	validKinds = []string{kindInt, kindDouble}
)

func (c *Number) splitMessageAndCollectParameters(ctx *appcontext.AppContext) ([]string, error) {
	span := sentryio.NewSpan(ctx.Context, "split message and collect parameters", "")
	defer span.Finish()

	parts := strings.Fields(c.Message)
	if len(parts) < 2 {
		return nil, errors.New("invalid command parameters")
	}
	return parts[2:], nil
}

func (c *Number) mapParametersToStruct(ctx *appcontext.AppContext, list []string) error {
	span := sentryio.NewSpan(ctx.Context, "map parameters to struct", "")
	defer span.Finish()

	var (
		parameters         = NumberParameters{}
		totalMatchedParams = 0
	)

	for _, opt := range list {
		// split
		parts := strings.Split(opt, "=")

		// if the option doesn't have 2 parts, it's invalid and just skip
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case appcommand.RandomParameters.Type:
			if !slices.Contains(validKinds, value) {
				continue
			}
			parameters.Kind = value
			totalMatchedParams++
		case appcommand.RandomParameters.Min:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Min = v
			totalMatchedParams++
		case appcommand.RandomParameters.Max:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Max = v
			totalMatchedParams++
		case appcommand.RandomParameters.Count:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Count = v
			totalMatchedParams++
		case appcommand.RandomParameters.Unique:
			v, err := strconv.ParseBool(value)
			if err != nil {
				continue
			}
			parameters.Unique = v
			totalMatchedParams++
		default:
			continue
		}
	}

	if totalMatchedParams > 0 {
		c.Parameters = &parameters
		return nil
	} else {
		return errors.New("invalid parameters")
	}
}

func (c *Number) random(ctx *appcontext.AppContext) (string, error) {
	span := sentryio.NewSpan(ctx.Context, "generate random number", "")
	defer span.Finish()

	if c.Parameters == nil {
		return "", errors.New("invalid parameters")
	}

	if c.Parameters.Min > c.Parameters.Max {
		return "", errors.New("min cannot be greater than max")
	}

	if c.Parameters.Count <= 0 {
		c.Parameters.Count = 1
	}

	if c.Parameters.Kind == "" {
		c.Parameters.Kind = kindInt // Default to "int" if not specified
	}

	// adjust count if unique is true and exceeds the possible unique values
	if c.Parameters.Kind == kindInt && c.Parameters.Unique && c.Parameters.Count > (c.Parameters.Max-c.Parameters.Min+1) {
		c.Parameters.Count = c.Parameters.Max - c.Parameters.Min + 1
	}

	if c.Parameters.Count > 100 {
		c.Parameters.Count = 100
	}

	var result strings.Builder
	result.WriteString("Result: \n")
	if c.Parameters.Kind == kindInt {
		values := map[int]bool{}
		for i := 0; i < c.Parameters.Count; i++ {
			n := c.randomInt(&values)
			result.WriteString(fmt.Sprintf("%d\\. `%d` \n", i+1, n))
		}
	} else {
		values := map[float64]bool{}
		for i := 0; i < c.Parameters.Count; i++ {
			n := c.randomDouble(&values)
			result.WriteString(fmt.Sprintf("%d\\. %.1f \n", i+1, n))
		}
	}

	return result.String(), nil
}

func (c *Number) randomInt(values *map[int]bool) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(c.Parameters.Max-c.Parameters.Min+1) + c.Parameters.Min
	if !c.Parameters.Unique {
		(*values)[n] = true
		return n
	}
	if _, exists := (*values)[n]; !exists {
		(*values)[n] = true
		return n
	} else {
		return c.randomInt(values)
	}
}

func (c *Number) randomDouble(values *map[float64]bool) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Float64()*(float64(c.Parameters.Max-c.Parameters.Min)) + float64(c.Parameters.Min)
	if !c.Parameters.Unique {
		(*values)[n] = true
		return n
	}
	if _, exists := (*values)[n]; !exists {
		(*values)[n] = true
		return n
	} else {
		return c.randomDouble(values)
	}
}

func (c *Number) Process(ctx *appcontext.AppContext) string {
	// collect parameters
	parameterList, err := c.splitMessageAndCollectParameters(ctx)
	if err != nil {
		return err.Error()
	}

	// map to struct
	if err = c.mapParametersToStruct(ctx, parameterList); err != nil {
		return err.Error()
	}

	// random
	result, err := c.random(ctx)
	if err != nil {
		return err.Error()
	}

	return result
}
