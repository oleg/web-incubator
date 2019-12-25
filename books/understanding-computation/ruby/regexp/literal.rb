require './pattern'

class Literal < Struct.new(:character)
  include Pattern
  
  def precedence
    3
  end

  def to_s
    character
  end

end
