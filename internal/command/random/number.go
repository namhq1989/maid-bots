package random

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/namhq1989/maid-bots/pkg/sentryio"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Number struct {
	Arguments  map[string]string
	Parameters *NumberParameters
}

type NumberParameters struct {
	Format string
	Min    int
	Max    int
	Count  int
	Unique bool
}

var (
	formatInt    = "int"
	formatDouble = "double"
	validFormats = []string{formatInt, formatDouble}
)

func (c *Number) mapParametersToStruct(ctx *appcontext.AppContext) error {
	span := sentryio.NewSpan(ctx.Context, "map parameters to struct")
	defer span.Finish()

	var (
		parameters         = NumberParameters{}
		totalMatchedParams = 0
	)

	for key, value := range c.Arguments {
		switch key {
		case appcommand.RandomNumberParameters.Format:
			if !slices.Contains(validFormats, value) {
				continue
			}
			parameters.Format = value
			totalMatchedParams++
		case appcommand.RandomNumberParameters.Min:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Min = v
			totalMatchedParams++
		case appcommand.RandomNumberParameters.Max:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Max = v
			totalMatchedParams++
		case appcommand.RandomNumberParameters.Count:
			v, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			parameters.Count = v
			totalMatchedParams++
		case appcommand.RandomNumberParameters.Unique:
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
	span := sentryio.NewSpan(ctx.Context, "generate random number")
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

	if c.Parameters.Format == "" {
		c.Parameters.Format = formatInt // Default to "int" if not specified
	}

	// adjust count if unique is true and exceeds the possible unique values
	if c.Parameters.Format == formatInt && c.Parameters.Unique && c.Parameters.Count > (c.Parameters.Max-c.Parameters.Min+1) {
		c.Parameters.Count = c.Parameters.Max - c.Parameters.Min + 1
	}

	if c.Parameters.Count > 100 {
		c.Parameters.Count = 100
	}

	var result strings.Builder
	result.Grow(c.Parameters.Count*10 + len("Result: \n"))
	result.WriteString("Result: \n")
	if c.Parameters.Format == formatInt {
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
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := seededRand.Intn(c.Parameters.Max-c.Parameters.Min+1) + c.Parameters.Min
	if !c.Parameters.Unique {
		(*values)[n] = true
		return n
	}
	for {
		if _, exists := (*values)[n]; !exists {
			(*values)[n] = true
			return n
		}
		n = seededRand.Intn(c.Parameters.Max-c.Parameters.Min+1) + c.Parameters.Min
	}
}

func (c *Number) randomDouble(values *map[float64]bool) float64 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := seededRand.Float64()*(float64(c.Parameters.Max-c.Parameters.Min)) + float64(c.Parameters.Min)
	if !c.Parameters.Unique {
		(*values)[n] = true
		return n
	}
	for {
		if _, exists := (*values)[n]; !exists {
			(*values)[n] = true
			return n
		}
		n = seededRand.Float64()*(float64(c.Parameters.Max-c.Parameters.Min)) + float64(c.Parameters.Min)
	}
}

func (c *Number) Process(ctx *appcontext.AppContext) string {
	// map to struct
	if err := c.mapParametersToStruct(ctx); err != nil {
		return err.Error()
	}

	// random
	result, err := c.random(ctx)
	if err != nil {
		return err.Error()
	}

	return result
}
