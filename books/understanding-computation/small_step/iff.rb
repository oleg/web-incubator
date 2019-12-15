class If < Struct.new(:condition, :consequence, :alternative)
  
  def to_s
    "if (#{condition}) { #{consequence} } else { #{alternative} }"
  end

  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end

  def reduce environment
    if condition.reducible?
      reduced = condition.reduce(environment)
      [If.new(reduced, consequence, alternative), environment]
    else
      case condition
      when Boolean.new(true)
        [consequence, environment]
      when Boolean.new(false)
        [alternative, environment]
      end
    end
  end
  
end
