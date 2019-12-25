require './pattern'

class Choose < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    0
  end
  
  def to_s
    #[first, second].map { |pattern| pattern.bracket(precedence) }.join('|')
    [first, second].join('|')
  end
  
end
