puts('Quantos alunos irá registrar?')
qtd_alunos = gets().to_i()
alunos = ''
for i in 1..qtd_alunos
  puts('Digite o nome do aluno')
  alunos = gets().to_s()
  puts('Digite as notas do aluno: ')
    nota = gets().to_f()
    nota2 = gets().to_f()
    media = (nota + nota2) / 2
    if media > 6 
      puts("A media do #{alunos} foi de #{media}, aprovado!")
    else
      puts("A media do #{alunos} foi de #{media}, reprovado!")
    end
end

