# # Create a single vaccine record
# vaccine = Vaccine.create(username: "Vaccine A")

# # # Create a single manufacturer record
# manufacturer = Manufacturer.create(username: "Manufacturer X")

# # # Create a single health facility record
# healthfacility = Healthfacility.create(username: "Health Facility 1")

# # # Create a few distributor records
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


# db/seeds.rb

Distributor.destroy_all
Healthfacility.destroy_all
Manufacturer.destroy_all

# Create seed data for distributors
distributors = Distributor.create!([
  {
    username: 'distributor1',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 1,
    manufacturer_id: 1,
    healthfacility_id: 1
  },
  {
    username: 'distributor2',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 2,
    manufacturer_id: 2,
    healthfacility_id: 2
  },
  {
    username: 'distributor3',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 3,
    manufacturer_id: 3,
    healthfacility_id: 3
  }
])

# Create seed data for health facilities
health_facilities = Healthfacility.create!([
  {
    username: 'healthfacility1',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 1,
    distributor_id: distributors.first.id
  },
  {
    username: 'healthfacility2',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 2,
    distributor_id: distributors.second.id
  },
  {
    username: 'healthfacility3',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 3,
    distributor_id: distributors.third.id
  }
])

# Create seed data for manufacturers
manufacturers = Manufacturer.create!([
  {
    username: 'manufacturer1',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 1,
    distributor_id: distributors.first.id
  },
  {
    username: 'manufacturer2',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 2,
    distributor_id: distributors.second.id
  },
  {
    username: 'manufacturer3',
    password_digest: BCrypt::Password.create("password123"),
    vaccine_id: 3,
    distributor_id: distributors.third.id
  }
])

puts "Seed data created successfully!"
