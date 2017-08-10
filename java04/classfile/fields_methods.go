package classfile

type MemberInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []*AttributeInfo
}

func readMember(reader *ClassReader) *MemberInfo {
	memberInfo := &MemberInfo{}
	memberInfo.readMemberInfo(reader)
	return memberInfo
}

func (self *MemberInfo) readMemberInfo(reader *ClassReader) {
	self.accessFlags = reader.readUint16()
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
	self.attributes = readAttributes(reader)
}
