import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Inventory } from './inventory.entity';

@Injectable()
export class InventoryService {
  constructor(
    @InjectRepository(Inventory)
    private inventoryRepository: Repository<Inventory>,
  ) {}

  async findAll(): Promise<Inventory[]> {
    return this.inventoryRepository.find();
  }

  async findOne(id: number): Promise<Inventory | null> {
    return this.inventoryRepository.findOneBy({ id });
  }

  async checkStock(productName: string): Promise<number> {
    const item = await this.inventoryRepository.findOneBy({ productName });
    return item ? item.quantity : 0;
  }

  async updateStock(productName: string, quantity: number): Promise<Inventory> {
    let item = await this.inventoryRepository.findOneBy({ productName });
    if (!item) {
      item = this.inventoryRepository.create({ productName, quantity });
    } else {
      item.quantity = quantity;
    }
    return this.inventoryRepository.save(item);
  }
}
