class Add < Struct.new(:left, :right)

  def evaluate environment
    l = left.evaluate(environment)
    r = right.evaluate(environment)
    Number.new(l.value + r.value)
  end
  
end
