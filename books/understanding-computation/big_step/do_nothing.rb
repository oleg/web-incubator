class DoNothing < Struct.new(:name, :expression)
  
  def evaluate environment
    environment
  end
  
end
