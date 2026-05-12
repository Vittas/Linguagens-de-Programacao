class Professor
  attr_accessor :nome_professor, :cod_funcionario, :disciplina
  attr_reader :ferias
  def initialize(nome_professor, cod_funcionario, disciplina)
    @nome_professor = nome_professor
    @cod_funcionario = cod_funcionario
    @disciplina = disciplina
    @ferias = false
    
  end

  def iniciar_ferias()
      @inicio_ferias =  Time.now()
      @ferias = true
  end

  def encerrar_ferias(
      @encerrar_ferias = Time.now()
      @ferias = false
  )
end