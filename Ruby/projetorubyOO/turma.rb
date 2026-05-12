class Turma
  attr_accessor :nome_turma, :numero_turma, :professor_turma
  def initialize(nome_turma, numero_turma, professor_turma)
    @nome_turma = nome_turma
    @numero_turma = numero_turma
    @professor_turma = professor_turma
  end
end
