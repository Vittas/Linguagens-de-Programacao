def Cal_notas()
  aluno = []
  x=0
  puts ("Quantos alunos ira registrar? ")
  n = gets().to_i

  for i in 1..n
    alunos_ar = [{ nome: "", nota: "", disciplina: ""}]
    puts ("Digite o nome do aluno: ")
    nome = gets().to_s.chomp()
    alunos_ar[0][:nome] = nome
    puts ("Digite a nota do aluno: ")
    nota = gets().to_s.chomp()
    alunos_ar[0][:nota] = nota
    puts ("Digite a disciplina da prova: ")
    disciplina = gets().to_s.chomp()
    alunos_ar[0][:disciplina] = disciplina
    aluno.concat(alunos_ar)
    puts("#{aluno[x][:nome]} tirou nota #{aluno[x][:nota]} na disciplina de #{aluno[x][:disciplina]}")
    x += 1
  end

  puts(aluno)
end

Cal_notas()