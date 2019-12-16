class Assign < Struct.new(:name, :expression)
  
  def to_ruby
    "-> e { e.merge({ #{name.inspect} => (#{expression.to_ruby})[e] }) }"
  end
  
end
