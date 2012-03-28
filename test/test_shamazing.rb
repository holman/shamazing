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

  def test_cli
    output = `bin/shamazing #{@sha}`.chomp

    assert_match /Longest string: cddfaedaea/, output
    assert_match /Longest integer: 20191041554/, output
  end

  def test_cli_string
    assert_equal 'cddfaedaea', `bin/shamazing #{@sha} -s`.chomp
    assert_equal 'cddfaedaea', `bin/shamazing #{@sha} --string`.chomp
  end

  def test_cli_integer
    assert_equal '20191041554', `bin/shamazing #{@sha} -i`.chomp
    assert_equal '20191041554', `bin/shamazing #{@sha} --integer`.chomp
  end
end