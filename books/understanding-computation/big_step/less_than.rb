class LessThan < Struct.new(:left, :right)

  def evaluate environment
    l = left.evaluate(environment)
    r = right.evaluate(environment)
    Boolean.new(l.value < r.value)
  end
  
end
