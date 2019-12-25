require './pattern'

class Repeat < Struct.new(:pattern)
  include Pattern

  def precedence
    2
  end
  
  def to_s
    pattern.bracket(precedence) + '*'
  end

end
