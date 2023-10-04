package soma

func Soma(valores ...int) (total int) {
    for _, valor := range valores {
        total += valor
    }

    return
}
