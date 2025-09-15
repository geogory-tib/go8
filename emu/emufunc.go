package emu

import (
	"bufio"
	"fmt"
	"go8/opcodes"
	"go8/types"
	"log"
	"os"
)

func Load_rom(filename string, chip8 *types.Chip8) {
	rom_file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(rom_file)
	buffer := make([]byte, 4026)
	_, err = reader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	for i := types.PROGRAM_ADDR; i < len(buffer); i++ {
		chip8.Ram[i] = buffer[i-types.PROGRAM_ADDR]
	}
	index := types.FONT_ADDR
	for _, font_byte := range types.Font_Set {
		chip8.Ram[index] = font_byte
		index++
	}
	chip8.Emu_state = types.RUNNING
	buffer = nil
	chip8.PC = types.PROGRAM_ADDR
}

func Chip8_cycle(chip8 *types.Chip8) {
	op := fetch_op(chip8)
	fmt.Printf("%x  | %b\n", op, op)
	decode_op(op, chip8)

}

func fetch_op(chip *types.Chip8) (op uint16) {
	var first_byte byte = chip.Ram[chip.PC]
	var second_byte byte = chip.Ram[chip.PC+1]
	chip.PC += 2
	temp_1 := uint16(first_byte)
	temp_2 := uint16(second_byte)
	op = temp_1<<8 | temp_2
	return
}

func decode_op(op uint16, chip *types.Chip8) {
	if op == 0xE0 {
		x := 0
		for i := range len(chip.Display) {
			chip.Display[i][0] = false
			x++
		}
		return
	}
	op_type := op >> 12
	op_type = op_type << 12
	switch op_type {
	case opcodes.DRAW:
		sprite_address := chip.I
		length := (op & 0x000F)
		x_reg := (op & 0x0F00) >> 8
		y_reg := (op & 0x00F0) >> 4
		x_pos := chip.V[x_reg]
		y_pos := chip.V[y_reg]
		for y := y_pos; y < byte(length); y++ {
			x_len := 0
			for x := x_len; x_len < 8; x++ {
				sprite_byte := chip.Ram[sprite_address]
				sprite_bit := (sprite_byte&(0x01<<x_pos))>>x_pos - 1
				if x >= 64 {
					break
				}
				if sprite_bit != 0 {
					chip.Display[y][x] = true
				}
				x++
			}
			y++
			if y >= 32 {
				break
			}
		}
	case opcodes.JUMP:
		address_shift := op << 12
		address := address_shift >> 12
		chip.PC = address
	case opcodes.ADD:
		reg_addres := (op & 0x0F00) >> 8
		add_value := (op & 0x00FF)
		chip.V[reg_addres] += byte(add_value)
	case opcodes.SET_I:
		adress_shift := op << 4
		chip.I = adress_shift >> 4
	case opcodes.SET:
		reg_addres := (op & 0x0F00) >> 8
		reg_value := (op & 0x00FF)
		chip.V[reg_addres] = byte(reg_value)
	}
}
