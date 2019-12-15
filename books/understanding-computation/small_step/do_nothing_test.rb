require './do_nothing'

require './test_setup'


class DoNothingTest < Test::Unit::TestCase
  def test_new
    dn = DoNothing.new
    assert_equal DoNothing.new, dn
  end

  def test_to_s
    assert_equal "do-nothing", DoNothing.new.to_s
  end

  def test_inspect
    assert_equal "«do-nothing»", DoNothing.new.inspect
  end
  
  def test_reducible
    assert_false DoNothing.new.reducible?
  end

end
