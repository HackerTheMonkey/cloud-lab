#!/bin/bash -euo pipefail

tag_name=code_matters_droplet_tag
api_key=$(cat ~/.digitalocean_api_key)

function delete_if_present(){
	local droplet_id=$(curl --silent \
	     --header "Authorization: Bearer ${api_key}" \
	     "https://api.digitalocean.com/v2/droplets?tag_name=${tag_name}"  | jq '.droplets[0] | .id')
	[[ "${droplet_id}" != "null" ]] && delete_droplet ${droplet_id} || echo "Droplet with tag ${tag_name} does not exist"
}

function delete_droplet(){
	echo "deleting droplet with tag ${tag_name}"
	local droplet_id=${1}
	curl -X DELETE \
	     --silent \
	     --header "Authorization: Bearer ${api_key}" \
	     "https://api.digitalocean.com/v2/droplets/${droplet_id}"
}

function get_ssh_key_id(){
	curl -X GET \
		--silent \
		--header "Content-Type: application/json" \
		--header "Authorization: Bearer ${api_key}" \
		"https://api.digitalocean.com/v2/account/keys" | jq '.ssh_keys[0] | .id'
}

function create_droplet(){	
	echo "creating a droplet, tagging as ${tag_name}"
	
	curl \
	 --silent \
	 --output output.log \
         --header "Authorization: Bearer ${api_key}" \
         --header "Content-Type: application/json" \
         "https://api.digitalocean.com/v2/droplets" \
         --data "{
		  \"name\": \"codematters-dryrun\",
		  \"region\": \"lon1\",
		  \"size\": \"s-1vcpu-1gb\",
		  \"image\": \"centos-8-x64\",
		  \"ssh_keys\": [
		    $(get_ssh_key_id)
		  ],
		  \"backups\": false,
		  \"ipv6\": false,
		  \"user_data\": \"$(cat config/cloud_config.yml)\",
		  \"private_networking\": null,
		  \"volumes\": null,
		  \"tags\": [
		    \"${tag_name}\"
		  ]
		}" && echo "Droplet created with tag ${tag_name}" || echo "error while creating droplet, check output.log for more details."

	# run the following command to check the progress of 
	# the instance provisioning within the droplet
	# tail -f /var/log/cloud-init-output.log 

	# cloud-config references:
	# https://www.digitalocean.com/community/tutorials/how-to-use-cloud-config-for-your-initial-server-setup#finished-product
	# https://www.digitalocean.com/community/tutorials/an-introduction-to-cloud-config-scripting
	# https://gist.github.com/c0psrul3/f2627de45f7d244afa48b0fe191a9ece
	# https://coreos.com/os/docs/latest/cloud-config.html#:~:text=The%20cloud%2Dconfig%20file%20uses,a%20shell%20script%20(advanced).
	# https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	# https://cloudinit.readthedocs.io/en/latest/topics/modules.html

}

# main
delete_if_present
create_droplet
