class LessThan < Struct.new(:left, :right)
  
  def to_ruby
    "-> e { (#{left.to_ruby})[e] < (#{right.to_ruby})[e] }"
  end
  
end
