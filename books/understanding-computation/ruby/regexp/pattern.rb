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

  def matches? str
    to_nfa_design.accepts? str
  end
  
end
