class DoNothing
  
  def to_ruby
    "-> e { e }"
  end

  def ==(other)
    other.instance_of?(DoNothing)
  end
  
end
