Gem::Specification.new do |s|
  s.specification_version = 2 if s.respond_to? :specification_version=
  s.required_rubygems_version = Gem::Requirement.new(">= 0") if s.respond_to? :required_rubygems_version=
  s.rubygems_version = '1.3.5'

  s.name              = 'shamazing'
  s.version           = '0.0.3'
  s.date              = '2012-03-27'
  s.rubyforge_project = 'shamazing'

  s.summary     = "Short description used in Gem listings."
  s.description = "A library to discover amazing things about a SHA1 hash "+
                  "(or MD5 or whatever). It's sha-mazing. Almost as shamazing "+
                  "as that pun."

  s.authors  = ["Zach Holman"]
  s.email    = 'zach@zachholman.com'
  s.homepage = 'https://github.com/holman/shamazing'

  s.require_paths = %w[lib]

  s.executables = ["shamazing"]

  # = MANIFEST =
  s.files = %w[
    README.md
    Rakefile
    bin/shamazing
    lib/shamazing.rb
    shamazing.gemspec
    test/test_shamazing.rb
  ]
  # = MANIFEST =

  s.test_files = s.files.select { |path| path =~ /^test\/test_.*\.rb/ }
end