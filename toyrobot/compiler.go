package toyrobot

import (
	"fmt"
	"github.com/danwhitford/toyrobot/belt"
)

//go:generate stringer -type=RobotType
type RobotType byte

const (
	T_INT RobotType = iota
	T_DIRECTION
	T_BOOL
)

//go:generate stringer -type=Instruction
type Instruction byte

const (
	OP_PUSH_VAL Instruction = iota
	OP_EXEC_WORD
)

type RobotCompiler struct {
	tokens *belt.Belt[Token]
}

func (r *RobotCompiler) Compile(input []Token) ([]byte, error) {
	r.tokens = belt.NewBelt[Token](input)

	instructions := make([]byte, 0)
	for r.tokens.HasNext() {
		token, err := r.tokens.GetNext()
		if err != nil {
			return nil, err
		}
		switch token.Type {
		case TOKEN_NUMBER:
			instructions = append(
				instructions,
				byte(OP_PUSH_VAL),
				byte(T_INT),
				byte(token.Value.(int)),
			)
		case TOKEN_DIRECTION:
			instructions = append(
				instructions,
				byte(OP_PUSH_VAL),
				byte(T_DIRECTION),
				byte(token.Value.(Direction)),
			)
		case TOKEN_WORD:
			tokenVal, ok := token.Value.(string)
			if !ok {
				return nil, fmt.Errorf("invalid token value '%v'", token.Value)
			}
			bytes := append([]byte(tokenVal), 0)
			instructions = append(
				instructions,
				byte(OP_EXEC_WORD),
			)
			instructions = append(
				instructions,
				bytes...,
			)
		default:
			return nil, fmt.Errorf("invalid instruction '%v'", token)
		}
	}

	return instructions, nil
}
