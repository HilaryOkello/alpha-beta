class SessionsController < ApplicationController
    def create
    distributor = Distributor.find_by(username: params[:username])
    if distributor&.authenticate(params[:password])
      session[:distributor_id] = distributor.id
      render json: distributor, status: :created
    else
      render json: { error: "Invalid username or password" }, status: :unauthorized
    end
  end

  def destroy
    session[:distributor_id] = nil
    head :no_content
  end

  def create
    healthfacility = Healthfacility.find_by(username: params[:username])
    if healthfacility&.authenticate(params[:password])
      session[:healthfacility_id] = healthfacility.id
      render json: healthfacility, status: :created
    else
      render json: { error: "Invalid username or password" }, status: :unauthorized
    end
  end

  def destroy
    session[:healthfacility_id] = nil
    head :no_content
  end

  def create
    manufacturer = Manufacturer.find_by(username: params[:username])
    if manufacturer&.authenticate(params[:password])
      session[:manufacturer_id] = manufacturer.id
      render json: manufacturer, status: :created
    else
      render json: { error: "Invalid username or password" }, status: :unauthorized
    end
  end

  def destroy
    session[:manufacturer_id] = nil
    head :no_content
  end
end
