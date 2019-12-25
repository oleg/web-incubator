require './pattern'

class Empty
  include Pattern

  def precedence
    3
  end
  
  def to_s
    ''
  end
  
end
