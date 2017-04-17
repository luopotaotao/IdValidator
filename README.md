# IdValidator
中国大陆用户身份证校验
# 主要内容
1. 基本格式正则校验
2. 地区校验
3. 生日范围校验
4. 校验和校验

以上步骤全部校验通过才算是校验通过,否则将返回相应错误信息

# 示例
    package idvalidator

    import (
        "testing"
    )

    func TestValidate(t *testing.T) {
        id:="510723198006202551"
        _,err:=Validate(id)
        if err!=nil{
            t.Error(err)
        }
    }

