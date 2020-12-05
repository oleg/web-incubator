class Variable < Struct.new(:name)

  def to_s
    "#{name}"
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end
  
  def reduce environment
    environment[name]
  end

end

