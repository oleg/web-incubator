class Machine < Struct.new(:statement, :environment)

  def run
    while statement.reducible?
      log
      step
    end
    log
  end

  def step
    self.statement, self.environment = statement.reduce(environment)
  end

  def log
    puts self
  end

  def to_s
    "#{statement}, #{environment}"
  end

end
