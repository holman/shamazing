require 'test/unit'
require 'lib/shamazing'

class TestShamazing < Test::Unit::TestCase
  def setup
    @sha = '9c2cddfaedaea9689a22e37aaa20191041554fe8'
    @shas = %w(
      fdb31214c2cca29e4f723ad676cddb043bd73986
      0c4b61fc2c5e7dd5566d42d0de1c431984899ddf
      9c2cddfaedaea9689a22e376aa20191041554fe8
      f1b4c270f6746cbfff99bbf0f5a2388f4e509943
    )
  end

  def test_string
    assert_equal 'cddfaedaea', Shamazing.string(@sha)
  end

  def test_string_downcase
    assert_equal 'cddfaedaea', Shamazing.string(@sha.upcase)
  end

  def test_string_from_array
    assert_equal 'cddfaedaea', Shamazing.string_from_array(@shas)
  end

  def test_integer
    assert_equal 20191041554, Shamazing.integer(@sha)
  end

  def test_integer_from_array
    assert_equal 20191041554, Shamazing.integer_from_array(@shas)
  end

  def test_repeating
    assert_equal 'aaa', Shamazing.repeating(@sha)
  end

  def test_repeating_from_array
    assert_equal 'fff', Shamazing.repeating_from_array(@shas)
  end

  def test_cli
    output = `bin/shamazing #{@sha}`.chomp

    assert_match /Longest string: cddfaedaea/, output
    assert_match /Longest integer: 20191041554/, output
    assert_match /Longest repeating: aaa/, output
  end

  def test_cli_string
    assert_equal 'cddfaedaea', `bin/shamazing #{@sha} -s`.chomp
    assert_equal 'cddfaedaea', `bin/shamazing #{@sha} --string`.chomp
  end

  def test_cli_integer
    assert_equal '20191041554', `bin/shamazing #{@sha} -i`.chomp
    assert_equal '20191041554', `bin/shamazing #{@sha} --integer`.chomp
  end

  def test_cli_repeating
    assert_equal 'aaa', `bin/shamazing #{@sha} -r`.chomp
    assert_equal 'aaa', `bin/shamazing #{@sha} --repeating`.chomp
  end

  def test_cli_git
    assert_match /Longest string: /, `bin/shamazing`.chomp
    assert_match /Longest integer: /, `bin/shamazing`.chomp
    assert_match /Longest repeating: /, `bin/shamazing`.chomp
  end
end