APP       := test
TARGET    := hello
MFLAGS    :=
DFLAGS    :=
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,/Users/hengzi/go}
$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/TarsCloud/TarsGo/tars/makefile.tars.gomod))
