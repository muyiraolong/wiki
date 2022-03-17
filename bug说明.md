#bug
##信息
save会删除原有hello.txt的信息
##原因
应该是save过程中，**未成功读取body**，在写入hello.txt时，覆盖了原有内容
#bug位置
应该是在saveHandler函数，或者是save方法中