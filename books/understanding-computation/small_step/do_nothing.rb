class DoNothing
  
  def ==(other)
    other.instance_of?(DoNothing)
  end
  
  def to_s
    "do-nothing"
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    false
  end
  
end

