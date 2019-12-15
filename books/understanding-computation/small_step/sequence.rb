class Sequence < Struct.new(:first, :second)

  def to_s
    "#{first}; #{second}"
  end

  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end

  def reduce environment
    case first
    when DoNothing.new
      [second, environment]
    else
      st, env = first.reduce(environment)
      [Sequence.new(st, second), env]
    end
  end

end

