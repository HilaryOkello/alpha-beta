# # Create a single vaccine record
# vaccine = Vaccine.create(username: "Vaccine A")

# # Create a single manufacturer record
# manufacturer = Manufacturer.create(username: "Manufacturer X")

# # Create a single health facility record
# healthfacility = Healthfacility.create(username: "Health Facility 1")

# # Create a few distributor records
# Distributor.create(
#   username: "distributor1",
#   password_digest: BCrypt::Password.create("password123"),
#   vaccine_id: vaccine.id,
#   manufacturer_id: manufacturer.id,
#   healthfacility_id: healthfacility.id
# )
# distributor = Distributor.find_by(username: "distributor1")
# if distributor
#     distributor.destroy
#     puts "Distributor 'distributor1' deleted."
# end