# zsh plugin for nik

__nik_update() {
        nik update `pwd`
}

j() {
        cd `nik get $@`
}

add-zsh-hook chpwd __nik_update
