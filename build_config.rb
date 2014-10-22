MRuby::Build.new do |conf|
  toolchain :gcc

  enable_debug

  conf.gem 'mrbgems/mruby-dir'
  conf.gem 'mrbgems/mruby-io'
  conf.gem 'mrbgems/mruby-mtest'

  conf.gembox 'default'
end
