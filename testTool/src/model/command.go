package model

type Command interface {
	Execute() string
}

type MacroCommand struct {
	Commands []Command
	result   string
}

func (self *MacroCommand) Execute() string {
	var result string
	for _, command := range self.Commands {
		result += command.Execute() + "\n"
	}

	return result
}

func (self *MacroCommand) Append(command Command) {
	// sliceを使用して、commandを追加していく
	self.Commands = append(self.Commands, command)
}

func (self *MacroCommand) Undo() {
	if len(self.Commands) != 0 {
		self.Commands = self.Commands[:len(self.Commands)-1]
	}
}

func (self *MacroCommand) Clear() {
	self.Commands = []Command{}
}
