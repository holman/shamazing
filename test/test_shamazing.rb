require 'test/unit'
require 'lib/shamazing'

class TestShamazing < Test::Unit::TestCase
  def setup
    @sha = '9c2cddfaedaea9689a22e376aa20191041554fe8'
  end

  def test_string
    assert_equal 'cddfaedaea', Shamazing.string(@sha)
  end

  def test_string_downcase
    assert_equal 'cddfaedaea', Shamazing.string(@sha.upcase)
  end

  def test_integer
    assert_equal 20191041554, Shamazing.integer(@sha)
  end
end