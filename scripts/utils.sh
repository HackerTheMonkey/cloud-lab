
# Intended to run manually and sourced into
# an iteractive shell.

api_key=$(cat ~/.digitalocean_api_key)

function do.list_all_distribution_images(){
	curl --silent \
	     --header "Authorization: Bearer ${api_key}" \
	     "https://api.digitalocean.com/v2/images?type=distribution" | jq
}

function do.list_all_domains(){
	local is_verbose=${1}
	local jq_filter=$([[ ${is_verbose} == "-v" ]] && echo '' || echo '.domains | .[] | .name')

	curl -X GET \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/domains" | jq ${jq_filter}
}

function do.get_domain(){
	local domain_name=${1}	
	curl -X GET \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/domains/${domain_name}" | jq
}

function do.create_domain(){
	local domain_name=${1}
	curl \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/domains/" \
		--data "{
				  \"name\": \"${domain_name}\"				  
				}" | jq
}

function do.delete_domain(){
	local domain_name=${1}
	curl -X DELETE \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/domains/${domain_name}" | jq
}

function do.get_droplet_by_id(){
	local tag_name=${1}
	curl \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/droplets?tag_name=${tag_name}" | jq
}
