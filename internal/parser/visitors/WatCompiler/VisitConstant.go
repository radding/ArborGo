package wast

import (
	// "encoding/binary"
	"fmt"
	"github.com/radding/ArborGo/internal/parser/ast"
	"strconv"
)

// VisitConstant visits the constant object
func (c *Compiler) VisitConstant(node *ast.Constant) (ast.VisitorMetaData, error) {
	// var value []byte
	tp := ""
	switch node.Type {
	case "STRINGVAL":
		return visitString(c, node)
	case "NUMBER":
		tp = "number"
		number, err := strconv.Atoi(node.Value)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		c.EmitFunc("i64.const %d", number)
	case "CHARVAL":
		tp = "char"
		c.EmitFunc("i32.const %d", rune(node.Raw[1]))
	case "FLOAT":
		tp = "float"
		number, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return ast.VisitorMetaData{}, err
		}
		c.EmitFunc("f64.const %f", number)
	default:
		return ast.VisitorMetaData{}, fmt.Errorf("encountered unknown constant")
	}
	return ast.VisitorMetaData{
		Location: "STACK",
		Types:    tp,
	}, nil
}

// Converts a number to a little endian
func littleEndian(number int32) uint32 {
	// LE := binary.LittleEndian
	// BE := binary.BigEndian
	// b := make([]byte, 4)
	// LE.PutUint32(b, uint32(number))
	// fmt.Printf("Value: %X", BE.Uint32(b))
	return uint32(number)
	// return BE.Uint32(b)
}

func visitString(c *Compiler, node *ast.Constant) (ast.VisitorMetaData, error) {
	place := c.getUniqueID("string", "begin")
	val := node.Value[1 : len(node.Value)-1]
	c.AddLocal(place, "i32")
	c.EmitFunc("i32.const %d", c.stackPointer)
	c.EmitFunc("set_local %s", place)
	c.EmitFunc("i32.const %d", c.stackPointer)
	c.EmitFunc("i32.const %d", littleEndian(int32(len(val))))
	c.EmitFunc("i32.store")
	for _, char := range val {
		c.stackPointer += 4
		c.EmitFunc("i32.const %d", c.stackPointer)
		c.EmitFunc("i32.const %d", littleEndian(int32(char)))

		c.EmitFunc("i32.store")
	}
	return ast.VisitorMetaData{
		Location: strconv.Itoa(c.dataSize),
		Types:    "string",
	}, nil
}
