class DoNothing
  
  def evaluate environment
    environment
  end

  def ==(other)
    other.instance_of?(DoNothing)
  end
   
end
