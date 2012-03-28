class Shamazing
  VERSION = '0.0.1'

  # Finds the longest continuous String in a sha.
  #
  # Returns a String.
  def self.string(sha)
    sha.
      downcase.
      scan(/[a-f]+/).
      sort{|a,b| b.length <=> a.length}.
      first
  end

  # Finds the longest continuous Integer in a sha.
  #
  # Returns an Integer.
  def self.integer(sha)
    sha.
      scan(/[0-9]+/).
      sort{|a,b| b.length <=> a.length}.
      first.to_i
  end
end