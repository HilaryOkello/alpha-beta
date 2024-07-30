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
end
