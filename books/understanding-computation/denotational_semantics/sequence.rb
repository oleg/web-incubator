class Sequence < Struct.new(:first, :second)

  def to_ruby
    "-> e { (#{second.to_ruby})[(#{first.to_ruby})[e]] }"
  end
  
end
