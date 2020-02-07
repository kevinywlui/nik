# zsh plugin for nik

__nik_update() {
        nik update `pwd`
}

j() {
        cd `nik get $1`
}

add-zsh-hook chpwd __nik_update
