class Assign < Struct.new(:name, :expression)
  
  def evaluate environment
    res = expression.evaluate(environment)
    environment.merge(name => res)
  end
  
end
