module Pattern
  def bracket outer_precedence
    if outer_precedence > precedence
      "(#{to_s})"
    else
      to_s
    end
  end
  
  def inspect
    "/#{self}/"
  end
end
