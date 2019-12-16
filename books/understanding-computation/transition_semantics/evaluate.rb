class Evaluate < Struct.new(:expression)
  def to_s
    "eval(#{expression})"
  end
  
  def inspect
    "«#{self}»"    
  end

  def reducible?
    true
  end

  def reduce environment
    if expression.reducible?
      [Evaluate.new(expression.reduce(environment)), environment]
    else
      [DoNothing.new, environment]
    end
  end
end
