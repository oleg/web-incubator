require './pattern'

class Concatenate < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    1
  end
  
  def to_s
    [first, second].map { |pattern| pattern.bracket(precedence) }.join
  end
  
end
