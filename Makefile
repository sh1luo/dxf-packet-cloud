init_api:
	hz new -I idl -idl idl/packet/packet.proto -module packet_cloud


update_api:
	hz update -I idl -idl idl/packet/packet.proto