require 'sinatra'
require 'sinatra/namespace'
require 'faraday'
require 'faraday_middleware'

histories = {}
histories[1] = {'id': 1, 'customer_id': 1, 'item_id': 1}
histories[2] = {'id': 2, 'customer_id': 2, 'item_id': 2}
histories[3] = {'id': 3, 'customer_id': 3, 'item_id': 3}

set :port, 8081
set :bind, '0.0.0.0'
items_url = ENV['ITEMS_URL'].nil? ? "http://localhost:8083" : ENV['ITEMS_URL']

get '/ping' do
  'pong'
end

namespace '/api' do
  get '/customers/:id' do |id|
    history = histories[id.to_i]
    halt(404, {message: 'NotFound'}.to_json) if history.nil?
    conn = Faraday.new(url: items_url + "/api/items/#{history[:item_id]}") do |faraday|
      faraday.response :json, :parser_options => { :symbolize_names => true }
      faraday.response :logger
      faraday.adapter Faraday.default_adapter
    end
    begin
      resp = conn.get
    rescue => e
      p e.message
      halt(404, {message: 'Failed to get item'}.to_json)
    end
    {id: history[:id], customer_id: history[:customer_id], item_name: resp.body[:name]}.to_json
  end
end
