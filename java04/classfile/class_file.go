package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32       // 魔数
	minorVersion uint16       // 次版本号
	majorVersion uint16       // 主版本号
	constantPool ConstantPool // 常量池
	accessFlags  uint16       //类访问标志
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()
	cf = &ClassFile{}
	cr := &ClassReader{classData}
	cf.read(cr)
	return cf, nil
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.readConstantPool(reader)
	self.readAccessFlags(reader)
	self.readThisClass(reader)
	self.readSuperClass(reader)
	self.readInterfaces(reader)
	self.readFieldsAndMethods(reader)
}

//读取魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	self.magic = magic
}

//读取和检查版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//读取常量池
func (self *ClassFile) readConstantPool(reader *ClassReader) {
	self.constantPool = readContantPool(reader)
}

//读取类访问标志
func (self *ClassFile) readAccessFlags(reader *ClassReader) {
	self.accessFlags = reader.readUint16()
}

//读取当前类
func (self *ClassFile) readThisClass(reader *ClassReader) {
	self.thisClass = reader.readUint16()
}

//读取superclass
func (self *ClassFile) readSuperClass(reader *ClassReader) {
	self.superClass = reader.readUint16()
}

//读取接口
func (self *ClassFile) readInterfaces(reader *ClassReader) {
	interfacesCount := int(reader.readUint16())
	interfaces := make([]uint16, interfacesCount)
	for i := 0; i < interfacesCount; i++ {
		interfaces[i] = reader.readUint16()
	}
	self.interfaces = interfaces
}

//读取方法和字段信息
func (self *ClassFile) readFieldsAndMethods(reader *ClassReader) {
	fieldsCount := int(reader.readUint16())
	fields := make([]*MemberInfo, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		fields[i] = readMember(reader)
	}
	self.fields = fields

	methodsCount := int(reader.readUint16())
	methods := make([]*MemberInfo, methodsCount)
	for i := 0; i < methodsCount; i++ {
		methods[i] = readMember(reader)
	}
	self.methods = methods
}
