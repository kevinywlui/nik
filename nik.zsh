# zsh plugin for nik

# hook to update score after each cd
__nik_update() {
        nik update `pwd`
}
add-zsh-hook chpwd __nik_update.

# single-letter alias to access nik
j() {
        if [[ $# == 0 ]] then
                `nik list -s`
        else
                cd `nik get $@`
        fi
}

