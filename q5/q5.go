package q5

import ("fmt" "strings")

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	var textosemvogal, textopontos string
	vogais := "aeiouAEIOU"
	for _, char := range s {
		if !strings.ContainsRune(vogais, char) {
			textosemvogal += string(char)
		}
	}
	for x := 0; x < len(textosemvogal); x++ {
		textopontos += "." + string(textosemvogal[x])
	}
	for i := 0; i < len(textopontos); i++ {
		if strings.ToUpper(string(textopontos[i])) == string(textopontos[i]) {
			textopontos = textopontos[:i] + strings.ToLower(string(textopontos[i])) + textopontos[i+1:]
		}
	}
	return textopontos
}
